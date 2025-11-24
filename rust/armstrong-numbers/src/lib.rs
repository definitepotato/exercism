pub fn is_armstrong_number(num: u32) -> bool {
    if num == 0 {
        return true;
    }

    let mut sum = 0;
    let mut my_num = num;
    let len_num = num.to_string().len() as u32;

    while my_num > 0 {
        let digit = my_num % 10;
        sum += digit.pow(len_num);
        my_num /= 10;
    }

    sum == num
}
