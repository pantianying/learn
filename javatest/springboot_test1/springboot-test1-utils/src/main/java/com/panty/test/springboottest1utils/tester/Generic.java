package com.panty.test.springboottest1utils.tester;

import java.util.ArrayList;
import java.util.List;



public class Generic {
    public static void GenericTest() {
        // 正确,类型错误,编译阶段就会报错
        List<String> arrayList = new ArrayList<String>();
        //错误
        // List arrayList = new ArrayList();
        arrayList.add("a");
        //arrayList.add(100);

        for (int i = 0; i < arrayList.size(); i++) {
            String temp=(String)arrayList.get(i);
            System.out.println("泛型测试："+temp);
        }
    }
}
