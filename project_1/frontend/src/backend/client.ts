import axios, {AxiosInstance} from "axios";
import {backend} from "../config/backend";
import {User} from "../types/user";

class BackendClient {
    client: AxiosInstance

    constructor(url: string) {
        this.client = axios.create({
            baseURL: url,
            timeout: 1000,
        })
    }

    async listUsers(): Promise<User[]> {
       const res = await this.client.get<User[]>("/users");

       return res.data;
    }

    async createUser(data: Omit<User, 'id'>): Promise<void> {
        console.log(data)
        await this.client.post("/users", data)
    }

    async updateUser(id: string, data: Partial<Omit<User, 'id'>>): Promise<void> {
        await this.client.put(`/users/${id}`, data)
    }

    async deleteUser(id: string): Promise<void> {
        await this.client.delete(`/users/${id}`)
    }
}

export const client = new BackendClient(backend.url)