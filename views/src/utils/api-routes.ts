import ky from 'ky';

export class APIRoutes {
    static readonly routes = {
        install: {
            check: ['./api/v1/install', 'GET'],
            insert: ['./api/v1/install', 'POST'],
        },
        school: {
            create: ['./api/v1/schools', 'POST'],
            list: ['./api/v1/schools', 'GET'],
        },
        users: {
            create: ['./api/v1/users', 'POST'],
            login: ['./api/v1/users/auth', 'POST'],
        },
    }

    static token = '';

    static async do<T>(endpoint: string, method: string, data?: Record<string, string> | undefined, isForm = false) {
        const payload = data ? (isForm ? (new URLSearchParams(data)).toString() : JSON.stringify(data)) : undefined;
        const client = ky.create({
            prefixUrl: import.meta.env.VITE_API_URL,
            headers: {
                Authorization: `Bearer ${APIRoutes.token}`,
            },
        });

        const response = await client(endpoint, {
                method,
                body: payload,
                headers: {
                    Authorization: `Bearer ${APIRoutes.token}`,
                    ...!isForm && method !== 'GET' ? {
                        'Content-Type': 'application/json',
                    } : {},
                },
            })
            .then(res => res.json<T>()).catch(() => undefined);

        return response;
    }
}