import { Box, Typography } from "@mui/material";
import Scoreboard from "../components/Scoreboard";

export default function ScoreboardPage() {
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
      <Scoreboard />
      <Box m={2}></Box>
    </Box>
  );
}
