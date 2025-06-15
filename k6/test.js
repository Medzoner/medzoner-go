import http from "k6/http";
import { check, sleep } from "k6";

// Test configuration
export const options = {
    thresholds: {
        // Assert that 99% of requests finish within 500ms.
        http_req_duration: ["p(99) < 500"],
    },
    stages: [
        { duration: "30s", target: 15 },
        { duration: "1m", target: 15 },
        { duration: "20s", target: 0 },
    ],
};

// Simulated user behavior
export default function () {
    let res = http.get("http://localhost:8080/");
    check(res, { "status was 200": (r) => r.status === 200 });
    sleep(1);
}
