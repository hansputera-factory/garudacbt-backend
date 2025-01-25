import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { installationFormSchema } from "@/validations/installationFormSchema";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { cn } from "@/lib/utils";
import { APIRoutes } from "@/lib/api-routes";

export default function InstallationIndexPage() {
    const form = useForm<z.infer<typeof installationFormSchema>>({
        resolver: zodResolver(installationFormSchema),
        defaultValues: {
            school: {
                app_name: 'GarudaCBTX',
            },
        },
        mode: 'all',
        shouldFocusError: true,
    });

    const generateShortCode = (input: string): string => 
        input
          .split(/\s+/)
          .map(word => /^\d+$/.test(word) ? word : word.toUpperCase())
          .join("")
          + Math.random().toString(36).slice(-4).toUpperCase();
    
    const onSubmit = async (values: z.infer<typeof installationFormSchema>) => {
        const response = await APIRoutes.do([
            APIRoutes.routes.install.insert,
            {
                school: {
                    ...values.school,
                    short_code: generateShortCode(values.school.name),
                },
                user: {
                    ...values.user,
                    email: values.school.email,
                },
            },
            false,
            {
                headers: {
                    'X-Authorized-Key': values.app_key,
                },
            }
        ]);

        console.log(response);
    }

    return (
        <React.Fragment>
            <div className="flex items-center justify-center min-h-screen m-6">
                <Card className="w-[800px]">
                    <CardHeader>
                        <CardTitle>
                            GarudaCBTX Basic Installation
                        </CardTitle>
                        <CardDescription>
                            Setup your first GarudaCBTX Application
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <Form {...form}>
                            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
                                <div className="grid lg:grid-cols-2 gap-4">
                                    {/* School name */}
                                    <FormField
                                        control={form.control}
                                        name="school.name"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    School Name
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="Your school name (e.g. SMA Negeri 3 Palu)" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Your registered school name in education ministry
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />

                                    {/* School National ID */}
                                    <FormField
                                        control={form.control}
                                        name="school.school_national_id"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    School National ID (NPSN)
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="Your school national id or in Indonesia usually called NPSN" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Your unique identifier school in education ministry (in Indonesia, usually we called it NPSN)
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                </div>

                                <div hidden>
                                    {/* School Short Code */}
                                    <FormField
                                        control={form.control}
                                        name="school.short_code"
                                        disabled
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    School Short Code
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="AUTOMATICALLY FILLED WHEN SUBMIT" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Unique identifier to identify your school data
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                </div>

                                <Separator style={{
                                    borderTop: '2px solid black',
                                }} />

                                {/* Address */}
                                <FormField
                                    control={form.control}
                                    name="school.address"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>
                                                School Address
                                            </FormLabel>
                                            <FormControl>
                                                <Input placeholder="Your school address (e.g. Jln. Dewi Sartika, No.104" {...field} />
                                            </FormControl>
                                            <FormDescription>
                                                Your current active school address
                                            </FormDescription>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />

                                {/* Longitude And Latitude */}
                                <div className="grid lg:grid-cols-2 gap-4">
                                    <FormField
                                        control={form.control}
                                        name="school.longitude"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    School Address - Longitude
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="School address longitude" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    School address longitude
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />

                                    <FormField
                                        control={form.control}
                                        name="school.latitude"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    School Address - Latitude
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="School address latitude" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    School address latitude
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                </div>

                                <Separator style={{
                                    borderTop: '2px solid black',
                                }} />

                                {/* Headmaster */}
                                <div className="grid lg:grid-cols-2 gap-4">
                                    <FormField
                                        control={form.control}
                                        name="school.headmaster_name"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Headmaster Name
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="Your school headmaster name (e.g. H.Idris Ade,S.Pd.,M.Si.)" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Your school headmaster full name
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                    <FormField
                                        control={form.control}
                                        name="school.headmaster_id"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Headmaster ID (NIP)
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="Your school headmaster id or NIP" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Your school headmaster id or in Indonesia we called it NIP
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                </div>

                                <Separator style={{
                                    borderTop: '2px solid black',
                                }} />
                                
                                <div className={cn('grid', 'lg:grid-cols-3', 'gap-4')}>
                                    {/* Website */}
                                    <FormField
                                        control={form.control}
                                        name="school.website"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    School Website
                                                </FormLabel>
                                                <FormControl>
                                                    <Input type="url" placeholder="Your school website if exists" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Your school website URL (if exists)
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />

                                    {/* Email */}
                                    <FormField
                                        control={form.control}
                                        name="school.email"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    School E-Mail
                                                </FormLabel>
                                                <FormControl>
                                                    <Input type="email" placeholder="Your school email" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Your school email
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                    {/* App Name */}
                                    <FormField
                                        control={form.control}
                                        name="school.app_name"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Application Name
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="Your CBT Application name" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    You may customize the application name if needed (default to GarudaCBTX)
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                </div>

                                
                                <Separator style={{
                                    borderTop: '2px solid black',
                                }} />

                                <div className={cn('grid', 'lg:grid-cols-3', 'gap-4')}>
                                    {/* App Key */}
                                    <FormField
                                        control={form.control}
                                        name="app_key"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Application Authorization Key
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="Please input your secret authorization key here" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Here is the secret key to grant confidential access
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />

                                    {/* User username */}
                                    <FormField
                                        control={form.control}
                                        name="user.username"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Administrator Username
                                                </FormLabel>
                                                <FormControl>
                                                    <Input placeholder="Please input your administrator username" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    Here you may custom your administrator username account
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />

                                    {/* User password */}
                                    <FormField
                                        control={form.control}
                                        name="user.password"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Administrator Password
                                                </FormLabel>
                                                <FormControl>
                                                    <Input type={'password'} placeholder="Please input your admin password account" {...field} />
                                                </FormControl>
                                                <FormDescription>
                                                    You may customize your administrator account password here
                                                </FormDescription>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                </div>

                                <Button type={'submit'}>
                                    Install
                                </Button>
                            </form>
                        </Form>
                    </CardContent>
                </Card>
            </div>
        </React.Fragment>
    )
}