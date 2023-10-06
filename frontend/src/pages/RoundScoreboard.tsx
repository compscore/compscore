import { Box, Typography } from "@mui/material";
import { useParams } from "react-router-dom";
import RoundScoreboard from "../components/RoundScoreboard";

export default function RoundScoreBoard() {
  const { round: round } = useParams() as { round: string };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography
        component='h1'
        variant='h3'
        fontWeight={700}
        sx={{
          marginTop: 5,
        }}
      >
        Scoreboard
      </Typography>
      <RoundScoreboard round={round} />
    </Box>
  );
}
