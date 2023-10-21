export type TeamScoreboard = {
  round: number;
  checks: [
    {
      name: string;
      status: [number];
    }
  ];
};
