package org.sonar.commonruleengine.checks;

import java.util.Arrays;
import java.util.List;

public class CheckList {

  public static List<Class<? extends Check>> getChecks() {
    return Arrays.asList(
      BinaryOperatorIdenticalExpressionsCheck.class,
      NoIdenticalFunctionsCheck.class,
      NoHardcodedCredentialsCheck.class,
      NoSelfAssignmentCheck.class,
      TooManyParametersCheck.class
    );
  }

}
