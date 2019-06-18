package com.panty.test1;


import org.springframework.stereotype.Component;

@Component
public class ArithmeticCalculatorImpl implements ArithmeticCalculator {
    @Override
    public int add(int i, int j) {
        return i + j;
    }
    @Override
    public int sub(int i, int j) {
        return i - j;
    }

}
