import { useParams } from "react-router-dom";

export default function TeamScoreboard() {
  const { team: number } = useParams();

  return <div>{number}</div>;
}
