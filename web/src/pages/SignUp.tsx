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
            <MDBValidationItem className="d-flex">
              <MDBInput
                value={formValue.username}
              />
            </MDBValidationItem>
            <MDBValidationItem>
              <MDBInput
                value={formValue.password}
              />
            </MDBValidationItem>
          </MDBValidation>
        </MDBCardBody>
      </MDBCard>
    </MDBContainer>
  );
}