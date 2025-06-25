import http from "k6/http";
import { check, sleep } from "k6";
import { Rate } from "k6/metrics";

export let errorRate = new Rate("errors");

// Config
export let options = {
  stages: [
    //  Low
    { duration: "20s", target: 20 }, // Ramp up to 20 users over 20 seconds
    //  Average
    { duration: "40s", target: 50 }, // Ramp up to 50 users over 40 seconds
    //  Peak
    { duration: "60s", target: 150 }, // Ramp up to 150 users over 60 seconds
    // Cool down
    { duration: "10s", target: 0 },
  ],
  thresholds: {
    http_req_duration: ["p(99)<1500"], // 99% of requests must complete below 1.5s
    http_req_failed: ["rate<0.1"], // Error rate must be below 10%
  },
};

const BASE_URL = "http://localhost:8080";

// Test data
const TEST_USER = {
  pin: "123456",
  user_id: "0befecd8-fccb-417e-aa0a-1a23c021f413",
};

export default function () {
  // Test 1: Health check
  let healthRes = http.get(`${BASE_URL}/health`);
  check(healthRes, {
    "health check status is 200": (r) => r.status === 200,
  }) || errorRate.add(1);

  // Test 2: Login
  let loginRes = http.post(
    `${BASE_URL}/api/auth/login`,
    JSON.stringify({
      pin: TEST_USER.pin,
      user_id: TEST_USER.user_id,
    }),
    {
      headers: { "Content-Type": "application/json" },
    }
  );

  check(loginRes, {
    "login status is 200": (r) => r.status === 200,
  }) || errorRate.add(1);

  let token = "";
  if (loginRes.status === 200) {
    try {
      token = JSON.parse(loginRes.body).data.access_token;
    } catch (e) {
      console.log("Failed to parse login response");
    }
  }

  // Test 3: Get user preview (public endpoint)
  let previewRes = http.get(
    `${BASE_URL}/api/users/${TEST_USER.user_id}/preview`
  );
  check(previewRes, {
    "user preview status is 200": (r) => r.status === 200,
  }) || errorRate.add(1);

  // Protected endpoints
  if (token) {
    const headers = {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    };

    // Test 4: Get me
    let meRes = http.get(`${BASE_URL}/api/users/me`, { headers });
    check(meRes, {
      "get me status is 200": (r) => r.status === 200,
    }) || errorRate.add(1);

    // Test 5: Get my banners
    let bannersRes = http.get(`${BASE_URL}/api/banners`, { headers });
    check(bannersRes, {
      "get banners status is 200": (r) => r.status === 200,
    }) || errorRate.add(1);

    // Test 6: Get my transactions
    let transactionsRes = http.get(`${BASE_URL}/api/transactions`, { headers });
    check(transactionsRes, {
      "get transactions status is 200": (r) => r.status === 200,
    }) || errorRate.add(1);

    // Test 7: Get my accounts
    let accountsRes = http.get(`${BASE_URL}/api/accounts`, { headers });
    check(accountsRes, {
      "get accounts status is 200": (r) => r.status === 200,
    }) || errorRate.add(1);
  }

  sleep(1); // think time
}
