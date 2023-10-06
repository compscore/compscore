import { Box } from "@mui/material";
import { useParams } from "react-router";
import Status from "../components/Status";

export default function StatusPage() {
  const { check: check, team: team } = useParams() as unknown as {
    check: string;
    team: number;
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        width: "80%",
        margin: "auto",
      }}
    >
      <Status check={check} team={team} />
    </Box>
  );
}
