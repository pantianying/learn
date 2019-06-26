package com.panty.test.springboottest1utils.dubbotest;

import com.alibaba.dubbo.config.ApplicationConfig;
import com.alibaba.dubbo.config.ReferenceConfig;
import com.alibaba.dubbo.config.RegistryConfig;
import com.alibaba.dubbo.rpc.RpcContext;
import com.alibaba.dubbo.rpc.service.GenericService;

public class DubboInvokeUtil {
  //为什么注入不了需要知道
//  @Resource private RegistryConfig registryConfig;
//  @Resource private ApplicationConfig applicationConfig;

  public Object invokeDubbo(Object[] parameterObjects) {
    ApplicationConfig applicationConfig = new ApplicationConfig();
    applicationConfig.setName("pantytest");
    RegistryConfig registryConfig = new RegistryConfig();
    registryConfig.setAddress("zookeeper://10.0.200.5:2181");
    registryConfig.setProtocol("zookeeper");
    applicationConfig.setRegistry(registryConfig);


    System.out.println("registryConfig:"+registryConfig);
    System.out.println("applicationConfig:"+applicationConfig);


    ReferenceConfig<GenericService> referenceConfig = new ReferenceConfig<>();
    referenceConfig.setApplication(applicationConfig);
    referenceConfig.setRegistry(registryConfig);

    referenceConfig.setInterface("com.tuya.aries.client.service.IFunctionMService");
    referenceConfig.setGeneric(true);
    GenericService genericService = referenceConfig.get();
    String[] paramTypes = new String[] {"xxxxx"};
    RpcContext.getContext().setAttachment("pantytest", "pantytest");
    Object returnValue = genericService.$invoke("queryDeviceStatus", paramTypes, parameterObjects);
    // void 返回为null
    return returnValue;
  }
}
