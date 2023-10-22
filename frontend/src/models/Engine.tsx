export type EngineStatus = {
  status: "error" | "running" | "paused" | "unknown";
  message: string;
};
