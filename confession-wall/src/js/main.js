// 确保在文档完全加载后再执行以下代码
document.addEventListener('DOMContentLoaded', function () {
    // 初始化 DataTable
    $("#example").DataTable({
        aaSorting: [],
        responsive: true,

        columnDefs: [
            {
                responsivePriority: 1,
                targets: 0
            },
            {
                responsivePriority: 2,
                targets: -1
            }
        ]
    });

    // 设置搜索框的占位符和样式
    $(".dataTables_filter input")
        .attr("placeholder", "Search here...")
        .css({
            width: "300px",
            display: "inline-block"
        });

    // 初始化 tooltip
    $('[data-toggle="tooltip"]').tooltip();
});
