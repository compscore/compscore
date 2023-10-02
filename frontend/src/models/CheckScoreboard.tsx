export type CheckScoreboard = {
  round: number;
  teams: [
    {
      name: string;
      status: [number];
    }
  ];
};
