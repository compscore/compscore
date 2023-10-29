export const domain = import.meta.env.DOMAIN;
export const api_url = import.meta.env.DEV ? "http://localhost:8080" : "";
export const path = "/";

export const short_refresh = 5000;
export const medium_refresh = 15000;
export const long_refresh = 30000;

export const default_timeout = 10000;

export const fetchWithTimeout = (
  url: RequestInfo,
  options?: RequestInit,
  timeout: number = default_timeout
): Promise<Response> => {
  return Promise.race([
    fetch(url, options),
    new Promise<Response>((_, reject) =>
      setTimeout(() => reject(new Error("Request timed out")), timeout)
    ),
  ]);
};
