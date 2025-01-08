import { describe, expect, it } from "bun:test"
import { faker } from '@faker-js/faker'

class MockData {
    static Username = faker.internet.username()
    static Password = faker.internet.password()
}

describe("SignUp test", () => {
    it("should return 201 for a SignUp", async () => {
        const response = await fetch("http://localhost:3432/auth/signup", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username: MockData.Username, password: MockData.Password })
        })
        expect(response.status).toBe(201)
    })
})