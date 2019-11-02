import React from "react";
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

export default function TableHeadCmp(props) {
    const fields = props.fields;
    return (
        <TableHead>
            <TableRow>
                { fields.map(function(item, index) {
                    return (<TableCell key={item.label}>{item.label}</TableCell>
                );
            })}
            </TableRow>
        </TableHead>
    );
}
