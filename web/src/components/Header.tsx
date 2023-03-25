import React, { useState } from 'react';
import {
  MDBNavbar,
  MDBNavbarNav,
  MDBNavbarItem,
  MDBNavbarLink,
  MDBNavbarToggler,
  MDBContainer,
  MDBIcon,
  MDBBtn,
  MDBCollapse
} from 'mdb-react-ui-kit';

export default function Header() {
    //const [showBasic, setShowBasic] = useState(false);

  return (
    <>
        <header>
            <MDBNavbar expand="lg" light bgColor='white' className="fixed-top">
                <MDBContainer fluid>
                    <MDBNavbarToggler
                        aria-controls='navbarExample01'
                        aria-expanded='false'
                        aria-label='Toggle navigation'
                    >
                        <MDBIcon fas icon='bars' />
                    </MDBNavbarToggler>

                    <MDBCollapse>
                        <MDBNavbarNav right className='mb-2 mb-lg-0'>
                            <MDBNavbarItem active>
                                <MDBNavbarLink aria-current='page' href='#'>
                                    Home
                                </MDBNavbarLink>
                            </MDBNavbarItem>
                            <MDBNavbarItem>
                                <MDBNavbarLink href='#'>Features</MDBNavbarLink>
                            </MDBNavbarItem>
                            <MDBNavbarItem>
                                <MDBNavbarLink href='#'>Pricing</MDBNavbarLink>
                            </MDBNavbarItem>
                            <MDBNavbarItem>
                                <MDBNavbarLink href='#'>About</MDBNavbarLink>
                            </MDBNavbarItem>
                        </MDBNavbarNav>
                    </MDBCollapse>
                </MDBContainer>
            </MDBNavbar>
        </header>
    </>
  );
}