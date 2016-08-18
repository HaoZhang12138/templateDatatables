var editor; // use a global for the submit and return data rendering in the examples
var table;

$(document).ready(function() {
    var tableName = $("table").attr("id");
    editor = new $.fn.dataTable.Editor( {
        ajax: "/defaultview?tableName=" + tableName,
        table: "#"+ tableName,
        idSrc: "courseid",
        fields: [{
            label: "CourseID:",
            name: "courseid",
            type: "hidden"
        },{
            label: "CourseName:",
            name: "coursename"
        }, {
            label: "TeacherName:",
            name: "teachername"
        }, {
            label: "Overview:",
            name: "overview"
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
            { data: "courseid"   },
            { data: "coursename"  },
            { data: "teachername"  },
            { data: "overview"  }
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
