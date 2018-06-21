// https://my.oschina.net/sannychan/blog/1537814
//使用go sdk获取连接
func Refresh_connect() {
	ctx, _ := context.WithCancel(context.Background())
	u, err := soap.ParseURL(cfg.Vsphere_url) # webservice地址，类似：https://10.50.82.155/sdk
	if err != nil {
		log.Fatal(err)
		return
	}
	u.User = url.UserPassword(cfg.Vsphere_user, cfg.Vsphere_passwd)
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		log.Fatal(err)
	}
	connect = &Connect{
		connect: c,
		context: ctx,
	}
}
//获取主机
func Get_hosts() ([]mo.HostSystem) {
	c := get_connect()
	client := c.connect
	m := view.NewManager(client.Client)
	v, err := m.CreateContainerView(c.context, client.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	log_process("create container view ")
	if err != nil {
		log.Fatal(err)
	}
	var hss []mo.HostSystem
	err = v.Retrieve(c.context, []string{"HostSystem"}, []string{"summary", "vm" , "datastore"}, &hss)
	if err != nil {
		log.Fatal(err)
	}
	log_process("retrieve host data ")
	return hss
}
//获取主机信息

//通过mo.HostSystem对象可以获取主机的cpu mhz速度、当前使用的mhz，最大内存、当前使用的内存信息，如果想要获取disk和net io，得通过查询监控数据才能获取。
host_cpu_speed := hs.Summary.Hardware.CpuMhz
cpu_capacity := float64(int64(hs.Summary.Hardware.CpuMhz) * int64(hs.Summary.Hardware.NumCpuCores))
cpu_usage := float64(hs.Summary.QuickStats.OverallCpuUsage)
mem_capacity := float64(hs.Summary.Hardware.MemorySize / (1024 * 1024))
mem_usage := float64(hs.Summary.QuickStats.OverallMemoryUsage)
//获取主机磁盘空间使用情况
//主机的磁盘信息主要是mo.Datastore对象，但hostSystem.Datastore获取到的是[]types.ManagedObjectReference，需要做一次转换

func Get_Host_disk_info(mos []types.ManagedObjectReference) (int64, int64, []*DiskInfo) {
	c := get_connect()
	pc := property.DefaultCollector(c.connect.Client)
	dss := []mo.Datastore{}
	err := pc.Retrieve(c.context, mos, []string{"summary"}, &dss)
	if err != nil {
		log.Error("get host datastore info error ", err)
	}
	var (
		capacity  int64
		free      int64
		diskInfos []*DiskInfo
	)
	for _, ds := range dss {
		capacity += ds.Summary.Capacity
		free += ds.Summary.FreeSpace
		diskInfos = append(diskInfos, &DiskInfo{
			Capacity: ds.Summary.Capacity / utils.GB,
			Free:     ds.Summary.FreeSpace / utils.GB,
			Path:     ds.Summary.Name,
		})
	}
	return capacity / utils.GB, free / utils.GB, diskInfos
}