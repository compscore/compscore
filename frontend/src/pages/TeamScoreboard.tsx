import { Box } from "@mui/material";
import { useParams } from "react-router-dom";
import TeamScoreboard from "../components/TeamScoreboard";

export default function TeamScoreboardPage() {
  const { team: team } = useParams() as unknown as { team: number };

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
