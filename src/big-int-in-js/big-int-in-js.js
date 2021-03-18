function test1(int1, int2) {
    var i = int1.length - 1;
    var j = int2.length - 1;
    var res = '';
    var flag = 0;
    while (i >= 0 || j >= 0||flag>=1) {
        var sum1 = flag;
        if (i >= 0 && j >= 0) {
            sum1 = Number(int1[i--]) + Number(int2[j--]) + sum1;
        }
        else if (i >= 0) {
            sum1 = Number(int1[i--]) + sum1;
        }
        else if (j >= 0) {
            sum1 = Number(int2[j--]) + sum1;
        }
        res=sum1%10+res
        flag = Math.floor(sum1 / 10);
    }
    return res;
}
test1('11111111111111111111111115', '999');
// 11111111111111111111112114