export type ResponseType<D> = {
    message: string;
    ok: boolean;
    data?: D;
}