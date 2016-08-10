var editor; // use a global for the submit and return data rendering in the examples
var table;

$(document).ready(function() {
    editor = new $.fn.dataTable.Editor( {
        ajax: "/view",
        table: "#example",
        idSrc: "_id",
        fields: [{
            label: "Id:",
            name: "_id",
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
            ],
        },{
            label: "Upload:",
            name: "fileid",
            type: "upload",
            display: function ( file_id ) {
                return ' <img src="' + table.file( 'files', file_id ).webpath + '"> </img> ';
            },
            clearText: "Clear",
            noImageText: 'No File'

            /*display: function ( id ) {

             return '<img src="'+editor.file( 'images', id ).webPath+'"/>';
             },
             noImageText: 'No image'
             /*  display: function () {
             return '<img height="20" width="20" src =http://pic.sc.chinaz.com/files/pic/pic9/201607/fpic5786.jpg' + '>';
             }*/

        }
        ]
    } );

    // Activate an inline edit on click of a table cell
   /* $('#example').on( 'click', 'tbody td:not(:first-child)', function (e) {
        editor.inline( this, {
            onBlur: 'submit'
        } );
    } );*/



   table =  $('#example').DataTable( {
        dom: "Bfrtip",
        ajax: {
            url: "/view"
        },
        columns: [
           /* {
                data: null,
                defaultContent: '',
                className: 'select-checkbox',
                orderable: false
            },*/
           /* {data: null, mRender: function () {
        return '<a href=/uploads/test.jpg' + '>';
    }},*/

            { data: "_id"   },
            { data: "user"  },
            { data: "pass"  },
            { data: "name"  },
            { data: "age"   },
            { data: "tel"   },
            { data: "sex"   },
            { data: "fileid",
                render: function ( file_id ) {
              //  alert(table.file( 'files', file_id ))
                    return file_id ?
                    ' <img src="' + table.file( 'files', file_id ).webpath + '"> </img>':
                    null;
                },
                defaultContent: "No File",
            }
        ],
        order: [ 1, 'asc' ],
        /*select: {
            style:    'os',
            selector: 'td:first-child'
        },*/
        select: true,
        buttons: [
            { extend: "create", editor: editor },
            { extend: "edit",   editor: editor },
            { extend: "remove", editor: editor }
        ]
    } );
} );