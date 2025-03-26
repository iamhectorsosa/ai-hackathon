"use client";

import { Skeleton } from "@/components/ui/skeleton";
import { api } from "@/lib/api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { toast } from "sonner";

const formSchema = z.object({
  message: z.string().min(2).max(50),
});

const FormDemo: React.FC = () => {
  const queryClient = useQueryClient();
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      message: "",
    },
  });

  const mutation = useMutation({
    mutationFn: api.greet,
    onSettled: () => form.reset(),
    onSuccess: (data) => {
      queryClient.resetQueries({ queryKey: api.getPosts.queryKey });
      queryClient.invalidateQueries({ queryKey: api.getStatus.queryKey });
      queryClient.invalidateQueries({ queryKey: api.getError.queryKey });
      toast.success(data.message);
    },
    onError: (error) => toast.error(error.message),
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    mutation.mutate({
      message: values.message,
    });
  }

  return (
    <div>
      <h2 className="mb-6 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors">
        Form
      </h2>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="message"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Message</FormLabel>
                <FormControl>
                  <Input placeholder="Hello World!" {...field} />
                </FormControl>
                <FormDescription>This is your message.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit">Submit</Button>
        </form>
      </Form>
    </div>
  );
};

const StatusDemo: React.FC = () => {
  const { status, data, error } = useQuery(api.getStatus);

  return (
    <div className="flex-1">
      <h2 className="mb-6 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors">
        Status
      </h2>
      {status === "pending" ? (
        <div className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200">
          <div className="p-6">
            <Skeleton className="h-6 w-3/4 mb-4" />
            <Skeleton className="h-4 w-full mb-2" />
          </div>
        </div>
      ) : status === "error" ? (
        <div className="bg-red-50 border-l-4 border-red-500 p-6 rounded shadow">
          <div className="ml-3">
            <h3 className="text-xl font-medium text-red-800 mb-2">
              Error loading status
            </h3>
            <p className="leading-7 text-red-700">
              {error?.message ?? "Unknown error"}
            </p>
          </div>
        </div>
      ) : (
        <div className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200 hover:shadow-lg transition-shadow duration-300">
          <div className="p-6">
            <h3 className="text-xl font-semibold tracking-tight mb-2">
              Status
            </h3>
            <p className="leading-7 text-gray-600">{JSON.stringify(data)}</p>
          </div>
        </div>
      )}
    </div>
  );
};

const ErrorDemo: React.FC = () => {
  const { status, data, error } = useQuery(api.getError);

  return (
    <div className="flex-1">
      <h2 className="mb-6 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors">
        Error
      </h2>
      {status === "pending" ? (
        <div className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200">
          <div className="p-6">
            <Skeleton className="h-6 w-3/4 mb-4" />
            <Skeleton className="h-4 w-full mb-2" />
          </div>
        </div>
      ) : status === "error" ? (
        <div className="bg-red-50 border-l-4 border-red-500 p-6 rounded shadow">
          <div className="ml-3">
            <h3 className="text-xl font-medium text-red-800 mb-2">
              Error loading error
            </h3>
            <p className="leading-7 text-red-700">
              {error?.message ?? "Unknown error"}
            </p>
          </div>
        </div>
      ) : (
        <div className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200 hover:shadow-lg transition-shadow duration-300">
          <div className="p-6">
            <h3 className="text-xl font-semibold tracking-tight mb-2">Data</h3>
            <p className="leading-7 text-gray-600">{JSON.stringify(data)}</p>
          </div>
        </div>
      )}
    </div>
  );
};

const Top3Posts: React.FC = () => {
  const { status, data, error } = useQuery(api.getPosts);
  return (
    <div className="mx-auto">
      <h2 className="mb-6 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors">
        Top 3 Posts
      </h2>
      {status === "pending" ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[1, 2, 3].map((i) => (
            <div
              key={i}
              className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200"
            >
              <div className="p-3">
                <Skeleton className="h-5 w-3/4 mb-2" />
                <Skeleton className="h-4 w-full" />
              </div>
            </div>
          ))}
        </div>
      ) : status === "error" ? (
        <div className="bg-red-50 border-l-4 border-red-500 p-4 rounded shadow">
          <div className="ml-3">
            <h3 className="text-sm font-medium text-red-800">
              Error loading posts
            </h3>
            <p className="mt-1 text-sm text-red-700">
              {error?.message ?? "Unknown error"}
            </p>
          </div>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {data.posts.slice(0, 3).map((post) => (
            <div
              key={post.id}
              className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200 hover:shadow-lg transition-shadow duration-300 flex items-center"
            >
              <div className="p-3 flex-grow">
                <h3 className="scroll-m-20 text-base font-semibold tracking-tight mb-1">
                  {post.title}
                </h3>
                <p className="text-sm text-gray-600 line-clamp-1">
                  {post.body}
                </p>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

const LatestsPosts: React.FC = () => {
  const { status, data, error } = useQuery(api.getPosts);

  return (
    <div className="mx-auto">
      <h2 className="mb-6 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors">
        Latest Posts
      </h2>
      {status === "pending" ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[1, 2, 3, 4, 5, 6].map((i) => (
            <div
              key={i}
              className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200"
            >
              <div className="p-6">
                <Skeleton className="h-6 w-3/4 mb-4" />
                <Skeleton className="h-4 w-full mb-2" />
                <Skeleton className="h-4 w-full mb-2" />
                <Skeleton className="h-4 w-2/3" />
              </div>
            </div>
          ))}
        </div>
      ) : status === "error" ? (
        <div className="bg-red-50 border-l-4 border-red-500 p-4 rounded shadow">
          <div className="ml-3">
            <h3 className="text-sm font-medium text-red-800">
              Error loading posts
            </h3>
            <p className="mt-1 text-sm text-red-700">
              {error?.message ?? "Unknown error"}
            </p>
          </div>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {data.posts.map((post) => (
            <div
              key={post.id}
              className="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200 hover:shadow-lg transition-shadow duration-300"
            >
              <div className="p-6">
                <h3 className="scroll-m-20 text-xl font-semibold tracking-tight mb-2">
                  {post.title}
                </h3>
                <p className="leading-7 text-gray-600">{post.body}</p>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default function Home() {
  return (
    <div className="space-y-12">
      <FormDemo />
      <div className="flex flex-col md:flex-row gap-6">
        <StatusDemo />
        <ErrorDemo />
      </div>
      <Top3Posts />
      <LatestsPosts />
    </div>
  );
}
