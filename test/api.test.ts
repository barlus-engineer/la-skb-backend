import { describe, expect, it } from "bun:test"
import { faker } from "@faker-js/faker"
import dotenv from "dotenv"

dotenv.config()

class MockData {
    static Username = faker.internet.username()
    static Password = faker.internet.password()
}

class TestUtils {
    static async postRequest(url: string, body: object) {
        const response = await fetch(url, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        })
        return response
    }

    static async assertStatus(url: string, body: object, expectedStatus: number) {
        const response = await this.postRequest(url, body)
        expect(response.status).toBe(expectedStatus)
    }
}

const IP = process.env.IP || "127.0.0.1"
const PORT = process.env.PORT || 3432

const BASE_URL = IP + ":" + PORT

describe("SignUp tests", () => {
    const endpoint = `${BASE_URL}/auth/signup`

    it("Incomplete form [No Username]", async () => {
        await TestUtils.assertStatus(endpoint, { username: "", password: MockData.Password }, 400)
    })

    it("Incomplete form [No Password]", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: "" }, 400)
    })

    it("Incomplete form [No Username and Password]", async () => {
        await TestUtils.assertStatus(endpoint, { username: "", password: "" }, 400)
    })

    it("Successful SignUp", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: MockData.Password }, 201)
    })

    it("Cannot SignUp again with the same Username", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: MockData.Password }, 400)
    })
})

describe("SignIn tests", () => {
    const endpoint = `${BASE_URL}/auth/signin`

    it("Incomplete form [No Username]", async () => {
        await TestUtils.assertStatus(endpoint, { username: "", password: MockData.Password }, 400)
    })

    it("Incomplete form [No Password]", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: "" }, 400)
    })

    it("Incomplete form [No Username and Password]", async () => {
        await TestUtils.assertStatus(endpoint, { username: "", password: "" }, 400)
    })

    it("Successful SignIn", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: MockData.Password }, 200)
    })

    it("Username not found", async () => {
        await TestUtils.assertStatus(endpoint, { username: ".", password: MockData.Password }, 400)
    })

    it("Incorrect Password", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: "." }, 400)
    })
})

describe("Delete Account tests", () => {
    const endpoint = `${BASE_URL}/auth/delete_account`

    it("Incomplete form [No Username]", async () => {
        await TestUtils.assertStatus(endpoint, { username: "", password: MockData.Password }, 400)
    })

    it("Incomplete form [No Password]", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: "" }, 400)
    })

    it("Incomplete form [No Username and Password]", async () => {
        await TestUtils.assertStatus(endpoint, { username: "", password: "" }, 400)
    })

    it("Successful Delete Account", async () => {
        await TestUtils.assertStatus(endpoint, { username: MockData.Username, password: MockData.Password }, 200)
    })
})
