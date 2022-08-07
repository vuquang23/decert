import { baseUrl, platform } from "const";

const request = async (
  method: "GET" | "POST" | "PUT" | "DELETE",
  path: string,
  content?: any,
  searchParams?: URLSearchParams
) => {
  if (content !== undefined && "platform" in content) {
    content.platform = platform.toString();
  }
  const response = await fetch(
    baseUrl +
      path +
      (searchParams !== undefined ? "?" + searchParams.toString() : ""),
    {
      method: method,
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(content),
    }
  );
  const json = await response.json();
  if (!response.ok || !("data" in json)) {
    throw Error(json.message ?? response.statusText);
  }
  return json.data;
};

const GET = (path: string, searchParams?: URLSearchParams) =>
  request("GET", path, undefined, searchParams);

const POST = (path: string, content: any) => request("POST", path, content);

const PUT = (path: string, content: any) => request("PUT", path, content);

const DELETE = (path: string) => request("DELETE", path);

export { GET, POST, PUT, DELETE };
