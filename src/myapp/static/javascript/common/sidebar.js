(function($) {

    function init() {

    }


    function bindEvent() {
        $('body')
        // 点击页面上的空白区域隐藏sidebar
        .on('click', '#main-content, #header', function(e) {

            // 判断点击的元素和其祖先元素是否是超链接，当不是超链接时才执行

            console.log($(e.target).closest('a')[0] !== undefined);
            if (!(e.target.tagName === 'A') && !($(e.target).closest('a')[0] !== undefined)) {
                var $sidebar = $('#sidebar');
                var $sidebarLeft = parseInt($sidebar.css('left'));
                var $pageWidth = parseInt($('body').css('width'));

                if ($pageWidth > 767) {
                    return false;
                }
                
                if ($sidebarLeft === 0) {

                    $('#sidebar').animate({
                        left: -250
                    }, 'normal');

                    $('#main-content').animate({
                        marginLeft: 0
                        // width: $pageWidth
                    }, 'normal');

                    $('#header').animate({
                        marginLeft: 0
                    }, 'normal');

                } else {
                    return false;
                }

            }   
        })
        // 显示隐藏sidebar
        .on('click', '#btn-open-sidebar', function() {
            var $sidebar = $('#sidebar');
            var $sidebarLeft = parseInt($sidebar.css('left'));
            var $pageWidth = parseInt($('body').css('width'));

            if ($pageWidth > 767) {
                return false;
            }
            
            if ($sidebarLeft === -250) {

                $('#sidebar').animate({
                    left: 0
                }, 'normal');

                $('#main-content').animate({
                    marginLeft: 250
                }, 'normal');

                $('#header').animate({
                    marginLeft: 250
                }, 'normal');

            } else if ($sidebarLeft === 0) {

                $('#sidebar').animate({
                    left: -250
                }, 'normal');

                $('#main-content').animate({
                    marginLeft: 0
                }, 'normal');

                $('#header').animate({
                    marginLeft: 0
                }, 'normal');

            } else {
                return false;
            }
        });

        // 调整页面尺寸时做一次把sidebar还原的操作(不这样做会导致布局异常)
        $(window).on('resize', function() {
            var $sidebar = $('#sidebar');
            var $sidebarLeft = parseInt($sidebar.css('left'));
            var $sidebarWidth = parseInt($sidebar.css('width'));
            var $pageWidth = parseInt($('body').css('width'));
            
            // 233是带滚动条的sidebar宽度
            if ($pageWidth < 767 && $sidebarLeft === 0 && ($sidebarWidth === 250 || $sidebarWidth === 233)) {
                console.log($sidebarWidth);
                $('#sidebar').animate({
                    left: -250
                }, 'normal');

                $('#main-content').animate({
                    marginLeft: 0
                }, 'normal');

                $('#header').animate({
                    marginLeft: 0
                }, 'normal');
            }
        });

    }

    init();
    bindEvent();
})(jQuery);