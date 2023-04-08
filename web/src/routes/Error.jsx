import { Navigate, useRouteError } from "react-router-dom";

export default function Error() {
  return (
    <Navigate to="/dashboard"/>
  );
}