import { z } from "zod";

export const installationFormSchema = z.object({
    name: z.string().min(10).max(255),
    short_code: z.string().optional(),
    school_national_id: z.string().min(8).refine((value) => /[0-9]+/gi.exec(value), 'Invalid school national id'),
    address: z.string().min(12).max(255),
    latitude: z.string().optional(),
    longitude: z.string().optional(),
    headmaster_name: z.string().min(3).max(100),
    headmaster_id: z.string().min(18).refine((value) => /[0-9]+/gi.exec(value), 'Invalid headmaster ID'),
    website: z.string().url({
        message: 'Invalid Website',
    }).optional(),
    email: z.string().email({
        message: 'Invalid email',
    }),
    app_name: z.string().min(5).default('GarudaCBTX'),
});
