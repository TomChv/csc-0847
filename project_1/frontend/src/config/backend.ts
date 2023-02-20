export interface Backend {
    url: string
}

const getEnv = (name: string): string => {
    const value = process.env[name]
    if (!value) {
        console.error(`${name} not found in environment`)
    }

    return value || ""
}

export const backend: Backend = {
    url: getEnv("REACT_APP_BACKEND_URL")
}