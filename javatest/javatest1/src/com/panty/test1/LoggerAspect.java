package com.panty.test1;

import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.springframework.stereotype.Component;
//放到容器中 声明为切

@Aspect
@Component
public class LoggerAspect {

    //声明是个前置通知
    @Before("execution(public int com.panty.test1.ArithmeticCalculatorImpl.add(int,int))")
    public  void beforeMethod(){
        System.out.println("before the method begin ");
    }
}
