import { useParams } from "react-router";
import {
  Box,
  Typography,
  TableContainer,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
} from "@mui/material";
import { StatusHistory } from "../models/StatusHistory";
import { useEffect, useState } from "react";

export default function Status() {
  const { check: check, team: team } = useParams() as unknown as {
    check: string;
    team: number;
  };

  const [data, setData] = useState<StatusHistory>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`http://localhost:8080/api/status/check/${check}/team/${team}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          let response = (await res.json()) as StatusHistory;
          if (res.status === 200) {
            setData(response);
          } else {
            console.log("Encountered an error");
          }
        })
        .catch((err) => {
          console.log("Encountered an error: " + err);
        });
    };

    fetchData();

    const pollingInterval = setInterval(fetchData, 1000);

    return () => clearInterval(pollingInterval);
  }, []);

  const [highlightedRound, setHighlightedRound] = useState<number | null>(null);

  const getBackgroundColor = (status: number, round: number) => {
    if (highlightedRound === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedRound === round) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (status === 0) {
      return "#f26359";
    } else if (status === 1) {
      return "#75ad77";
    }
    return "#cecdc6";
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
      <Typography
        component='h1'
        variant='h3'
        fontWeight={700}
        sx={{
          marginTop: 5,
        }}
        onClick={() => {
          window.location.href = "/scoreboard";
        }}
      >
        Team {team} - {check}
      </Typography>
      <Box m={2}></Box>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              {data?.map((status) => (
                <TableCell
                  size='small'
                  align='center'
                  key={"round-" + status.round}
                >
                  {status.round}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            <TableRow>
              {data?.map((status) => (
                <TableCell
                  size='medium'
                  align='center'
                  key={"status-" + status.round}
                  sx={{
                    backgroundColor: getBackgroundColor(
                      status.status,
                      status.round
                    ),
                  }}
                  onMouseEnter={() => {
                    setHighlightedRound(status.round);
                  }}
                  onMouseLeave={() => {
                    setHighlightedRound(null);
                  }}
                >
                  <Box height={75}></Box>
                </TableCell>
              ))}
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
      <Box m={2}></Box>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Round</Typography>
              </TableCell>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Time</Typography>
              </TableCell>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Status</Typography>
              </TableCell>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Error</Typography>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.map((status) => (
              <TableRow
                onMouseEnter={() => {
                  setHighlightedRound(status.round);
                }}
                onMouseLeave={() => {
                  setHighlightedRound(null);
                }}
                sx={{
                  backgroundColor:
                    highlightedRound === status.round
                      ? "#343434"
                      : "transparent",
                }}
              >
                <TableCell size='small' align='center'>
                  {status.round}
                </TableCell>
                <TableCell size='small' align='center'>
                  {status.time}
                </TableCell>
                <TableCell size='small' align='center'>
                  {status.status === 0
                    ? "Down"
                    : status.status === 1
                    ? "Up"
                    : "Unknown"}
                </TableCell>
                <TableCell size='small' align='center'>
                  {status.error}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
