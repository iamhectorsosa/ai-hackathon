import { queryOptions } from "@tanstack/react-query";

async function throwServerError(response: Response): Promise<void> {
  const errorData: ServerError = await response
    .json()
    .catch(() => ({ error: "Unknown error" }));
  throw new Error(
    errorData.error || `Request failed with status ${response.status}`,
  );
}

export const api = {
  getStatus: queryOptions({
    queryKey: ["getStatus"],
    queryFn: async (): Promise<Status> => {
      const response = await fetch("/_api/status");
      if (!response.ok) await throwServerError(response);
      return await response.json();
    },
  }),
  getError: queryOptions({
    queryKey: ["getError"],
    retry: 1,
    queryFn: async (): Promise<Status> => {
      const response = await fetch("/_api/error");
      if (!response.ok) await throwServerError(response);
      return await response.json();
    },
  }),
  getPosts: queryOptions({
    queryKey: ["getPosts"],
    queryFn: async (): Promise<Array<Post>> => {
      await new Promise((resolve) => setTimeout(resolve, Math.random() * 2000));
      const response = await fetch("/_api/posts");
      if (!response.ok) await throwServerError(response);
      return await response.json();
    },
  }),
  greet: async (payload: GreetPayload): Promise<GreetResponse> => {
    const response = await fetch("/_api/greet", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });
    if (!response.ok) await throwServerError(response);
    return await response.json();
  },
};

export type Status = {
  status: string;
};

export type Post = {
  id: number;
  title: string;
  body: string;
};

export type GreetPayload = {
  message: string;
};

export type GreetResponse = {
  message: string;
};

export type ServerError = {
  error: string;
};
