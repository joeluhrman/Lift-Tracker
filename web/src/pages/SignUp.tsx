import React, { useState } from 'react';
import {
  MDBContainer,
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
    <MDBContainer fluid>
      <MDBCard>
        <MDBCardTitle>Sign Up</MDBCardTitle>
        <MDBCardBody>

          <MDBValidation>

            <MDBValidationItem className="mb-3" feedback="Please choose a username." invalid>
              <MDBInput
                name="username"
                type="text"
                value={formValue.username}
                label="Username"
                required
                onChange={handleChange}
              />
            </MDBValidationItem>
            <MDBValidationItem className="mb-3" feedback="Please enter a valid email address." invalid>
              <MDBInput
                name="email"
                type="email"
                value={formValue.email}
                label="Email"
                required
                onChange={handleChange}
              />
            </MDBValidationItem>
            <MDBValidationItem className="mb-3" feedback="Please enter a password." invalid>
              <MDBInput
                name="password"
                type="password"
                value={formValue.password}
                label="Password"
                required
                onChange={handleChange}
              />
            </MDBValidationItem>

            <MDBBtn type="submit" className="col-12">Sign Up</MDBBtn>

          </MDBValidation>

        </MDBCardBody>
      </MDBCard>
    </MDBContainer>
  );
}