(function ($) {
    'use strict';
    $('.item').on("click", function () {
        $(this).next().slideToggle(100);
        // 移除了关闭其他元素的代码
    });
}(jQuery));
