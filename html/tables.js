var editor; // use a global for the submit and return data rendering in the examples
var table;

$(document).ready(function() {
    var tableName = $("table").attr("id");
    editor = new $.fn.dataTable.Editor( {
        ajax: "/defaultview?tableName=" + tableName,
        table: "#"+ tableName,
        idSrc: "id",
        fields: [{
            label: "Id:",
            name: "id",
            type: "hidden"
        },{
            label: "User:",
            name: "user"
        }, {
            label: "Pass:",
            name: "pass"
        }, {
            label: "Name:",
            name: "name"
        }, {
            label: "Age:",
            name: "age"
        },{
            label: "Tel:",
            name: "tel"
        },{
            label: "Sex:",
            name: "sex",
            type: "select",
            options: [
                { label: "男", value: "男" },
                { label: "女",  value: "女" }
            ]
        },{
            label: "Upload:",
            name: "fileid",
            type: "upload",
            display: function ( file_id ) {
                return ' <img src="' + table.file( 'files', file_id ).webpath + '"> </img> ';
            },
            clearText: "Clear",
            noImageText: 'No File'

        }
        ]
    } );

    table =  $("#"+ tableName).DataTable( {
        dom: "Bfrtip",
        pageLength: 8,
        ajax: {
            url: "/defaultview?tableName=" + tableName
        },
        columns: [
            { data: "id"   },
            { data: "user"  },
            { data: "pass"  },
            { data: "name"  },
            { data: "age"   },
            { data: "tel"   },
            { data: "sex"   },
            { data: "fileid",
                render: function ( file_id ) {
                    return file_id ?
                    ' <img src="' + table.file( 'files', file_id ).webpath + '"> </img>':
                        null;
                },
                defaultContent: "No File",
            }
        ],
        order: [ 1, 'asc' ],
        select: true,
        buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]
    } );
} );