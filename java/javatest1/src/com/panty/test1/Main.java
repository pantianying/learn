package com.panty.test1;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {
    public static void main(String[] args) {

        ApplicationContext ctx = new ClassPathXmlApplicationContext("spring-config.xml");

        ArithmeticCalculator atri = null;

        atri = ctx.getBean(ArithmeticCalculator.class);
        System.out.println("获取到的类："+atri.getClass().getName());
        int result = atri.add(1, 2);
        System.out.println("hello world!" + result);


    }
}
