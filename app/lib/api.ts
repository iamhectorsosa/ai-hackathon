import { queryOptions } from "@tanstack/react-query";
import { fetchGet, fetchPost } from "./api.helpers";

export const api = {
  getStatus: queryOptions({
    queryKey: ["getStatus"],
    queryFn: () => fetchGet<StatusReturn>("status"),
  }),
  getError: queryOptions({
    queryKey: ["getError"],
    retry: 1,
    queryFn: () => fetchGet<ErrorResponse>("error"),
  }),
  getPosts: queryOptions({
    queryKey: ["getPosts"],
    queryFn: () => fetchGet<PostReturn>("posts"),
  }),
  greet: (args: GreetArgs) => fetchPost<GreetArgs, GreetReturn>("greet", args),
};

export type StatusReturn = {
  status: string;
};

export type Post = {
  id: number;
  title: string;
  body: string;
};

export type PostReturn = {
  posts: Array<Post>;
};

export type GreetArgs = {
  message: string;
};

export type GreetReturn = {
  message: string;
};

export type ErrorResponse = {
  error: string;
};
