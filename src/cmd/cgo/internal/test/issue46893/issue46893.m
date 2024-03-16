#import "Panicker.objc.h"

void issue46893() {
    PanickerInit();

    PanickerPanickerImpl *p = [[PanickerPanickerImpl alloc] init];
    NSMutableData *d = [NSMutableData dataWithLength:10];
    memset(d.mutableBytes, 1, d.length);

    for (int i = 0; i < 10000000; i++) {
        NSError *error = nil;
        [p panic:d error:&error];
        if (error) {
            NSLog(@"Error: %@", error);
            break;
        }
    }
}
