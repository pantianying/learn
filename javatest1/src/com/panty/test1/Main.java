package com.panty.test1;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {
    public static void main(String[] args) {

        ApplicationContext ctx = new ClassPathXmlApplicationContext("spring-config.xml");

        ArithmeticCalculatorImpl atri = null;

//        String[] beanList = ctx.getBeanDefinitionNames();
//        for (String name:beanList){
//            System.out.println(name);
//        }
        atri = ctx.getBean(ArithmeticCalculatorImpl.class);
        int result = atri.add(1, 2);
        System.out.println("hello world!" + result);


    }
}
