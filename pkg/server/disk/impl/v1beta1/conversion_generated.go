// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta1"
	impl "github.com/kubernetes-csi/csi-proxy/pkg/server/disk/impl"
)

func autoConvert_v1beta1_DiskIDs_To_impl_DiskIDs(in *v1beta1.DiskIDs, out *impl.DiskIDs) error {
	return nil
}

// Convert_v1beta1_DiskIDs_To_impl_DiskIDs is an autogenerated conversion function.
func Convert_v1beta1_DiskIDs_To_impl_DiskIDs(in *v1beta1.DiskIDs, out *impl.DiskIDs) error {
	return autoConvert_v1beta1_DiskIDs_To_impl_DiskIDs(in, out)
}

func autoConvert_impl_DiskIDs_To_v1beta1_DiskIDs(in *impl.DiskIDs, out *v1beta1.DiskIDs) error {
	return nil
}

// Convert_impl_DiskIDs_To_v1beta1_DiskIDs is an autogenerated conversion function.
func Convert_impl_DiskIDs_To_v1beta1_DiskIDs(in *impl.DiskIDs, out *v1beta1.DiskIDs) error {
	return autoConvert_impl_DiskIDs_To_v1beta1_DiskIDs(in, out)
}

func autoConvert_v1beta1_DiskLocation_To_impl_DiskLocation(in *v1beta1.DiskLocation, out *impl.DiskLocation) error {
	out.Adapter = in.Adapter
	out.Bus = in.Bus
	out.Target = in.Target
	out.LUNID = in.LUNID
	return nil
}

// Convert_v1beta1_DiskLocation_To_impl_DiskLocation is an autogenerated conversion function.
func Convert_v1beta1_DiskLocation_To_impl_DiskLocation(in *v1beta1.DiskLocation, out *impl.DiskLocation) error {
	return autoConvert_v1beta1_DiskLocation_To_impl_DiskLocation(in, out)
}

func autoConvert_impl_DiskLocation_To_v1beta1_DiskLocation(in *impl.DiskLocation, out *v1beta1.DiskLocation) error {
	out.Adapter = in.Adapter
	out.Bus = in.Bus
	out.Target = in.Target
	out.LUNID = in.LUNID
	return nil
}

// Convert_impl_DiskLocation_To_v1beta1_DiskLocation is an autogenerated conversion function.
func Convert_impl_DiskLocation_To_v1beta1_DiskLocation(in *impl.DiskLocation, out *v1beta1.DiskLocation) error {
	return autoConvert_impl_DiskLocation_To_v1beta1_DiskLocation(in, out)
}

func autoConvert_v1beta1_DiskStatsRequest_To_impl_DiskStatsRequest(in *v1beta1.DiskStatsRequest, out *impl.DiskStatsRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_v1beta1_DiskStatsRequest_To_impl_DiskStatsRequest is an autogenerated conversion function.
func Convert_v1beta1_DiskStatsRequest_To_impl_DiskStatsRequest(in *v1beta1.DiskStatsRequest, out *impl.DiskStatsRequest) error {
	return autoConvert_v1beta1_DiskStatsRequest_To_impl_DiskStatsRequest(in, out)
}

func autoConvert_impl_DiskStatsRequest_To_v1beta1_DiskStatsRequest(in *impl.DiskStatsRequest, out *v1beta1.DiskStatsRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_impl_DiskStatsRequest_To_v1beta1_DiskStatsRequest is an autogenerated conversion function.
func Convert_impl_DiskStatsRequest_To_v1beta1_DiskStatsRequest(in *impl.DiskStatsRequest, out *v1beta1.DiskStatsRequest) error {
	return autoConvert_impl_DiskStatsRequest_To_v1beta1_DiskStatsRequest(in, out)
}

func autoConvert_v1beta1_DiskStatsResponse_To_impl_DiskStatsResponse(in *v1beta1.DiskStatsResponse, out *impl.DiskStatsResponse) error {
	out.DiskSize = in.DiskSize
	return nil
}

// Convert_v1beta1_DiskStatsResponse_To_impl_DiskStatsResponse is an autogenerated conversion function.
func Convert_v1beta1_DiskStatsResponse_To_impl_DiskStatsResponse(in *v1beta1.DiskStatsResponse, out *impl.DiskStatsResponse) error {
	return autoConvert_v1beta1_DiskStatsResponse_To_impl_DiskStatsResponse(in, out)
}

func autoConvert_impl_DiskStatsResponse_To_v1beta1_DiskStatsResponse(in *impl.DiskStatsResponse, out *v1beta1.DiskStatsResponse) error {
	out.DiskSize = in.DiskSize
	return nil
}

// Convert_impl_DiskStatsResponse_To_v1beta1_DiskStatsResponse is an autogenerated conversion function.
func Convert_impl_DiskStatsResponse_To_v1beta1_DiskStatsResponse(in *impl.DiskStatsResponse, out *v1beta1.DiskStatsResponse) error {
	return autoConvert_impl_DiskStatsResponse_To_v1beta1_DiskStatsResponse(in, out)
}

func autoConvert_v1beta1_ListDiskIDsRequest_To_impl_ListDiskIDsRequest(in *v1beta1.ListDiskIDsRequest, out *impl.ListDiskIDsRequest) error {
	return nil
}

// Convert_v1beta1_ListDiskIDsRequest_To_impl_ListDiskIDsRequest is an autogenerated conversion function.
func Convert_v1beta1_ListDiskIDsRequest_To_impl_ListDiskIDsRequest(in *v1beta1.ListDiskIDsRequest, out *impl.ListDiskIDsRequest) error {
	return autoConvert_v1beta1_ListDiskIDsRequest_To_impl_ListDiskIDsRequest(in, out)
}

func autoConvert_impl_ListDiskIDsRequest_To_v1beta1_ListDiskIDsRequest(in *impl.ListDiskIDsRequest, out *v1beta1.ListDiskIDsRequest) error {
	return nil
}

// Convert_impl_ListDiskIDsRequest_To_v1beta1_ListDiskIDsRequest is an autogenerated conversion function.
func Convert_impl_ListDiskIDsRequest_To_v1beta1_ListDiskIDsRequest(in *impl.ListDiskIDsRequest, out *v1beta1.ListDiskIDsRequest) error {
	return autoConvert_impl_ListDiskIDsRequest_To_v1beta1_ListDiskIDsRequest(in, out)
}

// detected external conversion function
// Convert_v1beta1_ListDiskIDsResponse_To_impl_ListDiskIDsResponse(in *v1beta1.ListDiskIDsResponse, out *impl.ListDiskIDsResponse) error
// skipping generation of the auto function

// detected external conversion function
// Convert_impl_ListDiskIDsResponse_To_v1beta1_ListDiskIDsResponse(in *impl.ListDiskIDsResponse, out *v1beta1.ListDiskIDsResponse) error
// skipping generation of the auto function

func autoConvert_v1beta1_ListDiskLocationsRequest_To_impl_ListDiskLocationsRequest(in *v1beta1.ListDiskLocationsRequest, out *impl.ListDiskLocationsRequest) error {
	return nil
}

// Convert_v1beta1_ListDiskLocationsRequest_To_impl_ListDiskLocationsRequest is an autogenerated conversion function.
func Convert_v1beta1_ListDiskLocationsRequest_To_impl_ListDiskLocationsRequest(in *v1beta1.ListDiskLocationsRequest, out *impl.ListDiskLocationsRequest) error {
	return autoConvert_v1beta1_ListDiskLocationsRequest_To_impl_ListDiskLocationsRequest(in, out)
}

func autoConvert_impl_ListDiskLocationsRequest_To_v1beta1_ListDiskLocationsRequest(in *impl.ListDiskLocationsRequest, out *v1beta1.ListDiskLocationsRequest) error {
	return nil
}

// Convert_impl_ListDiskLocationsRequest_To_v1beta1_ListDiskLocationsRequest is an autogenerated conversion function.
func Convert_impl_ListDiskLocationsRequest_To_v1beta1_ListDiskLocationsRequest(in *impl.ListDiskLocationsRequest, out *v1beta1.ListDiskLocationsRequest) error {
	return autoConvert_impl_ListDiskLocationsRequest_To_v1beta1_ListDiskLocationsRequest(in, out)
}

// detected external conversion function
// Convert_v1beta1_ListDiskLocationsResponse_To_impl_ListDiskLocationsResponse(in *v1beta1.ListDiskLocationsResponse, out *impl.ListDiskLocationsResponse) error
// skipping generation of the auto function

// detected external conversion function
// Convert_impl_ListDiskLocationsResponse_To_v1beta1_ListDiskLocationsResponse(in *impl.ListDiskLocationsResponse, out *v1beta1.ListDiskLocationsResponse) error
// skipping generation of the auto function

// detected external conversion function
// Convert_v1beta1_PartitionDiskRequest_To_impl_PartitionDiskRequest(in *v1beta1.PartitionDiskRequest, out *impl.PartitionDiskRequest) error
// skipping generation of the auto function

func autoConvert_impl_PartitionDiskRequest_To_v1beta1_PartitionDiskRequest(in *impl.PartitionDiskRequest, out *v1beta1.PartitionDiskRequest) error {
	return nil
}

// Convert_impl_PartitionDiskRequest_To_v1beta1_PartitionDiskRequest is an autogenerated conversion function.
func Convert_impl_PartitionDiskRequest_To_v1beta1_PartitionDiskRequest(in *impl.PartitionDiskRequest, out *v1beta1.PartitionDiskRequest) error {
	return autoConvert_impl_PartitionDiskRequest_To_v1beta1_PartitionDiskRequest(in, out)
}

func autoConvert_v1beta1_PartitionDiskResponse_To_impl_PartitionDiskResponse(in *v1beta1.PartitionDiskResponse, out *impl.PartitionDiskResponse) error {
	return nil
}

// Convert_v1beta1_PartitionDiskResponse_To_impl_PartitionDiskResponse is an autogenerated conversion function.
func Convert_v1beta1_PartitionDiskResponse_To_impl_PartitionDiskResponse(in *v1beta1.PartitionDiskResponse, out *impl.PartitionDiskResponse) error {
	return autoConvert_v1beta1_PartitionDiskResponse_To_impl_PartitionDiskResponse(in, out)
}

func autoConvert_impl_PartitionDiskResponse_To_v1beta1_PartitionDiskResponse(in *impl.PartitionDiskResponse, out *v1beta1.PartitionDiskResponse) error {
	return nil
}

// Convert_impl_PartitionDiskResponse_To_v1beta1_PartitionDiskResponse is an autogenerated conversion function.
func Convert_impl_PartitionDiskResponse_To_v1beta1_PartitionDiskResponse(in *impl.PartitionDiskResponse, out *v1beta1.PartitionDiskResponse) error {
	return autoConvert_impl_PartitionDiskResponse_To_v1beta1_PartitionDiskResponse(in, out)
}

func autoConvert_v1beta1_RescanRequest_To_impl_RescanRequest(in *v1beta1.RescanRequest, out *impl.RescanRequest) error {
	return nil
}

// Convert_v1beta1_RescanRequest_To_impl_RescanRequest is an autogenerated conversion function.
func Convert_v1beta1_RescanRequest_To_impl_RescanRequest(in *v1beta1.RescanRequest, out *impl.RescanRequest) error {
	return autoConvert_v1beta1_RescanRequest_To_impl_RescanRequest(in, out)
}

func autoConvert_impl_RescanRequest_To_v1beta1_RescanRequest(in *impl.RescanRequest, out *v1beta1.RescanRequest) error {
	return nil
}

// Convert_impl_RescanRequest_To_v1beta1_RescanRequest is an autogenerated conversion function.
func Convert_impl_RescanRequest_To_v1beta1_RescanRequest(in *impl.RescanRequest, out *v1beta1.RescanRequest) error {
	return autoConvert_impl_RescanRequest_To_v1beta1_RescanRequest(in, out)
}

func autoConvert_v1beta1_RescanResponse_To_impl_RescanResponse(in *v1beta1.RescanResponse, out *impl.RescanResponse) error {
	return nil
}

// Convert_v1beta1_RescanResponse_To_impl_RescanResponse is an autogenerated conversion function.
func Convert_v1beta1_RescanResponse_To_impl_RescanResponse(in *v1beta1.RescanResponse, out *impl.RescanResponse) error {
	return autoConvert_v1beta1_RescanResponse_To_impl_RescanResponse(in, out)
}

func autoConvert_impl_RescanResponse_To_v1beta1_RescanResponse(in *impl.RescanResponse, out *v1beta1.RescanResponse) error {
	return nil
}

// Convert_impl_RescanResponse_To_v1beta1_RescanResponse is an autogenerated conversion function.
func Convert_impl_RescanResponse_To_v1beta1_RescanResponse(in *impl.RescanResponse, out *v1beta1.RescanResponse) error {
	return autoConvert_impl_RescanResponse_To_v1beta1_RescanResponse(in, out)
}
