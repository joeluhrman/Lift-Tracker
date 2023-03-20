import React, { useState } from 'react';
import {
  MDBValidation,
  MDBValidationItem,
  MDBInput,
  MDBInputGroup,
  MDBBtn,
  MDBCheckbox,
  MDBCard,
  MDBCardBody,
  MDBCardHeader,
  MDBCardTitle,
} from 'mdb-react-ui-kit';

export default function SignUp() {
  const [formValue, setFormValue] = useState({
    username: "",
    email: "",
    password: "",
  })

  const handleChange = (e: any) => {
    setFormValue({ ...formValue, [e.target.name]: e.target.value });
  };

  return (
    <MDBCard>
      <MDBCardTitle>Sign Up</MDBCardTitle>
      <MDBValidation className="row g-3">
        <MDBValidationItem className="row-md-4">
          <MDBInput
            value={formValue.username}
            name="username"
            onChange={handleChange}
            id="username"
            required
            label="Username"
          />
        </MDBValidationItem>
        <MDBValidationItem className="row-md-4">
          <MDBInput
            value={formValue.password}
            name="password"
            onChange={handleChange}
            id="password"
            required
            label="Password"
          />
        </MDBValidationItem>
      </MDBValidation>
    </MDBCard>
  );
}