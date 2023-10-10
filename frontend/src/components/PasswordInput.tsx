import { useState } from "react";
import TextField from "@mui/material/TextField";
import IconButton from "@mui/material/IconButton";
import InputAdornment from "@mui/material/InputAdornment";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";

type props = {
  value: string;
  onBlur?: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>;
};

function PasswordInput({ value, onBlur }: props) {
  console.log(value);
  const [password, setPassword] = useState<string>(value);
  const [showPassword, setShowPassword] = useState(false);

  const handleTogglePassword = () => {
    setShowPassword((prevShowPassword) => !prevShowPassword);
  };

  return (
    <TextField
      type={showPassword ? "text" : "password"}
      variant='standard'
      value={password}
      onChange={(e) => {
        setPassword(e.target.value);
      }}
      onBlur={onBlur}
      fullWidth
      InputProps={{
        endAdornment: (
          <InputAdornment position='end'>
            <IconButton onClick={handleTogglePassword}>
              {showPassword ? <VisibilityOffIcon /> : <VisibilityIcon />}
            </IconButton>
          </InputAdornment>
        ),
      }}
    />
  );
}

export default PasswordInput;
