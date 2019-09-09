  /*
  function fnWrap(){ 
    
    function fnTouzi(){
        alert('亲，请关注我们的新产品！');
    }

    fnTouzi();

};

fnWrap();
*/

// 上面的写法还可能重名，可以改写成下面封闭函数的形式：
// 一个分号表示一个空的js语句，是合法的。
;;;;;;

// 在代码前面加分号，是为了在代码压缩成一行时，前面的代码在最后没有加分号，导致代码出错！
/*;(function(){ 
    
    function fnTouzi(){
        alert('亲，请关注我们的新产品！');
    }
    fnTouzi();

})();
*/


// 封闭函数装高手的写法：

/*;!function(){ 
    
    function fnTouzi(){
        alert('亲，请关注我们的新产品！');
    }
    fnTouzi();

}();
*/

;~function(){ 
    
    function fnTouzi(){
        alert('亲，请关注我们的新产品！');
    }
    fnTouzi();

}();


