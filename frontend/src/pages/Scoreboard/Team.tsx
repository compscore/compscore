import { Box } from "@mui/material";
import { useParams } from "react-router-dom";
import TeamScoreboard from "../../components/Scoreboard/Team";

export default function TeamScoreboardPage() {
  const { team: team } = useParams() as unknown as { team: string };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <TeamScoreboard team={team} />
    </Box>
  );
}
