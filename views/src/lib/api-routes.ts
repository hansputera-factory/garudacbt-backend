export type DoParams = [string[], any | undefined, boolean, RequestInit?];

export class APIRoutes {
    static readonly routes = {
        install: {
            check: ['/api/v1/install', 'GET'],
            insert: ['/api/v1/install', 'POST'],
        },
        school: {
            create: ['/api/v1/schools', 'POST'],
            list: ['/api/v1/schools', 'GET'],
        },
        users: {
            create: ['/api/v1/users', 'POST'],
            login: ['/api/v1/users/auth', 'POST'],
        },
    }

    static token = '';

    static async do(params: DoParams) {
        const [endpoint, data, isForm] = params;
        const payload = data ? (isForm ? (new URLSearchParams(data as Record<string, string>)).toString() : JSON.stringify(data)) : undefined;

        const response = await fetch(endpoint[0], {
            method: endpoint[1],
            body: payload,
            headers: {
                Authorization: `Bearer ${APIRoutes.token}`,
                ...!isForm && endpoint[1] !== 'GET' ? {
                    'Content-Type': 'application/json',
                } : {},
                ...params[3]?.headers,
            },
        }).then(res => res.json()).catch(() => undefined);

        return response;
    }
}