import axios, {AxiosInstance} from "axios";
import {Picture} from "../types/pictures";

class BackendClient {
    URL = "https://api-bfkpu45v2a-ue.a.run.app/"

    client: AxiosInstance

    constructor() {
        this.client = axios.create({
            baseURL: this.URL,
            timeout: 10000,
            withCredentials: false,
        })
    }

    async listPictures(): Promise<Picture[]> {
        const res = await this.client.get<{ files: Picture[] }>("/pictures")

        return res.data.files.map((file) => {
            let label = file.label.split(",")[0]
            if (label !== "Dog" && label !== "Person" && label !== "Flower") {
                label = "Other"
            }

            return ({...file, label })
        }) || []
    }

    async addPicture(file: File, metadata: Pick<Picture, "location" | "date" | "author">) {
        const formData = new FormData()

        formData.append("file", file)
        formData.append("metadata", JSON.stringify(metadata))

        await this.client.postForm("/pictures", formData)
    }

    async editPicture(name: string, metadata: Partial<Pick<Picture, "location" | "date" | "label" | "author">>): Promise<void> {
        const formData = new FormData()

        formData.append("metadata", JSON.stringify(metadata))

        await this.client.putForm(`/pictures/${name}`, formData)
    }

    async deletePicture(name: string): Promise<void> {
        await this.client.delete(`/pictures/${name}`)
    }
}

const client = new BackendClient();

export default client;