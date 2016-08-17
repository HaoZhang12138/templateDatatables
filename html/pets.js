var editor; // use a global for the submit and return data rendering in the examples
var table;

$(document).ready(function() {
    var tableName = $("table").attr("id");
    editor = new $.fn.dataTable.Editor( {
        ajax: "/view?tableName=" + tableName,
        table: "#"+ tableName,
        idSrc: "petid",
        fields: [{
            label: "PetId:",
            name: "petid",
            type: "hidden"
        },{
            label: "Name:",
            name: "name"
        }, {
            label: "Category:",
            name: "category"
        }, {
            label: "Color:",
            name: "color"
        }, {
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
            url: "/view?tableName=" + tableName,
            error: function () {
                alert("error")
            }
        },
        columns: [
            { data: "petid"   },
            { data: "name"  },
            { data: "category"  },
            { data: "color"  },
            { data: "fileid",
                render: function ( file_id ) {
                    return file_id ?
                    ' <img src="' + table.file( 'files', file_id ).webpath + '"> </img>':
                        null;
                },
                defaultContent: "No File"
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