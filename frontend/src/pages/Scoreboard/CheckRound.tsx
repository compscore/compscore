import { Box } from "@mui/material";
import { useParams } from "react-router-dom";
import CheckRoundScoreboard from "../../components/Scoreboard/CheckRound";

export default function CheckScoreboardPage() {
  const { check: check, round: round } = useParams() as {
    check: string;
    round: string;
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <CheckRoundScoreboard check={check} round={round} />
    </Box>
  );
}
