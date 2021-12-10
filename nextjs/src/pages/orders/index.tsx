// @flow 
import { GetServerSideProps } from 'next';
import * as React from 'react';
import axios from 'axios'
import Link from "next/link";
import { Typography, Link as MuiLink } from "@mui/material";
import { DataGrid, GridColumns } from '@mui/x-data-grid';
import { OrderStatus, OrderStatusTranslate } from '../../utils/models';
import { withIronSessionSsr } from "iron-session/next";
import ironConfig from '../../utils/iron-config';


const OrdersPage = (props: any) => {
  const columns: GridColumns = [
    {
      field: "id",
      headerName: "ID",
      width: 300,
      renderCell: (params) => {
        return (
          <Link href={`/orders/${params.value}`} passHref>
            <MuiLink>{params.value}</MuiLink>
          </Link>
        );
      },
    },
    {
      field: "amount",
      headerName: "Amount",
      width: 100,
    },
    {
      field: "credit_card_number",
      headerName: "Credit Card Num.",
      width: 200,
    },
    {
      field: "credit_card_name",
      headerName: "Credit Card Name",
      width: 200,
    },
    {
      field: "status",
      headerName: "Status",
      width: 110,
      valueFormatter: (params) =>
        OrderStatusTranslate[params.value as OrderStatus],
    },
  ];

  return (
    <div style={{ height: 400, width: "100%" }} >
      <Typography component="h1" variant="h4">My orders</Typography>
      <DataGrid
        columns={columns}
        rows={props.orders}
        rowsPerPageOptions={[5]}
        checkboxSelection
        disableSelectionOnClick
      />
    </div>
  );
};

export default OrdersPage;

export const getServerSideProps: GetServerSideProps = withIronSessionSsr(
  async (context) => {
    const account = context.req.session.account;

    if (!account) {
      return {
        redirect: {
          destination: "/login",
          permanent: false,
        },
      };
    }

    const { data } = await axios.get(
      `http://localhost:3001/api/orders`,
      {
        headers: {
          cookie: context.req.headers.cookie as string,
        },
      }
    );

    return {
      props: {
        orders: data,
      },
    };
  },
  ironConfig
);