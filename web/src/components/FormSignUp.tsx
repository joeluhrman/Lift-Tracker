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
import axios from "axios"

export default function FormSignUp() {
  const [formValue, setFormValue] = useState({
    username: "",
    email: "",
    password: "",
  })

  const handleChange = (e: any) => {
    setFormValue({ ...formValue, [e.target.name]: e.target.value });
  };

  const handleSubmit = async() => {
    var data = {
        username: formValue.username,
        email: formValue.email,
        password: formValue.password
    }

    try {
       var res = await axios.post("/api/v1/user", JSON.stringify(data))
       console.log(res)
    } catch(error) {
        console.log(error)
    }
  }

  return (
    <MDBContainer fluid>
      <MDBCard className="w-25">
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

            <MDBBtn type="submit" className="col-12" onClick={handleSubmit}>Sign Up</MDBBtn>

          </MDBValidation>

        </MDBCardBody>
      </MDBCard>
    </MDBContainer>
  );
}