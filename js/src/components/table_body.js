import React from "react";
import TableBody from "@material-ui/core/TableBody";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";

export default function TableBodyCmp(props) {
    return (
        <TableBody>
            {props.items.map((item, idx) => {
                return (
                    <TableRow key={idx}>
                         {props.fields.map((field, fidx) => {
                            var value = item[field.id];
                            return (
                                 <TableCell>{value}</TableCell>
                            )
                         })}                       
                    </TableRow>
                )
            })}
        </TableBody>
    )
}
