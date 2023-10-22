export type EngineStatus = {
  status: "error" | "running" | "paused" | "unknown";
  message: string;
};

export type EngineMessage = {
  message: string;
};
