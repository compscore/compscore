import { useParams } from "react-router";

export default function StatusHistory() {
  const { check: check, team: team } = useParams() as {
    check: string;
    team: number;
  };

  return (
    <>
      <h1>{check}</h1>
      <h1>{team}</h1>
    </>
  );
}
